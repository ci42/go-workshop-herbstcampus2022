package metasearch

import (
	"fmt"
	"github.com/ci42/ImgScraper/scraper"
	"github.com/ci42/go-workshop-herbstcampus2022/stats"
	"sync"
	"time"
)

func MetaImageSearch(q string, timeout time.Duration) ([]scraper.Image, stats.Stats, error) {
	var wg sync.WaitGroup
	imgs := []scraper.Image{}
	stats := stats.Stats{}

	finished := make(chan bool)
	image := make(chan scraper.Image)
	errC := make(chan error)

	getImages := func(imgScraper scraper.ImgScraper) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			imgs, err := imgScraper.Search(q)
			if err != nil {
				errC <- err
				return
			}
			for _, img := range imgs {
				image <- img
				//<-time.After(time.Duration(rand.Int31n(100)) * time.Millisecond)
			}
		}()
	}

	getImages(scraper.NewPexelsImgSearch(timeout, PexelsAPIKey))
	getImages(scraper.NewPixabayImgSearch(timeout, PixabayAPIKey))
	getImages(scraper.NewUnsplashImgSearch(timeout, UnsplashAPIKey))

	go func() {
		wg.Wait()
		finished <- true
	}()

	for {
		select {
		case <-time.After(timeout):
			return imgs, stats, fmt.Errorf("Timeout überschritten -- breche Suche ab.")
		case newImage := <-image:
			imgs = append(imgs, newImage)
			stats.Inc(newImage.Source)
		case err := <-errC:
			return imgs, stats, err
		case <-finished:
			return imgs, stats, nil
		}
	}

	return nil, nil, fmt.Errorf("uhm ... something weird happened")
}
