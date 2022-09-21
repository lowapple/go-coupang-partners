package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/fetch"
	"github.com/chromedp/cdproto/network"
	"log"
	"strings"
)
import "go-coupang-partners/src/common/config"
import "github.com/chromedp/chromedp"
import "github.com/jessevdk/go-flags"

func init() {
	_, err := flags.Parse(&config.Opts)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Client
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("window-size", "1280,720"),
		chromedp.Flag("no-sandbox", true),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	urlMap := map[network.RequestID]string{}
	// API Injection
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventLoadingFinished:
			url := urlMap[ev.RequestID]
			delete(urlMap, ev.RequestID)
			if strings.Contains(url, `search`) {
				go func(ev *network.EventLoadingFinished) {
					c := chromedp.FromContext(ctx)
					rbp := network.GetResponseBody(ev.RequestID)
					body, err := rbp.Do(cdp.WithExecutor(ctx, c.Target))
					if err != nil {
						fmt.Println(err)
					}
					if err == nil {
						fmt.Printf("%s\n", body)
					}
				}(ev)
			}
			break
		case *network.EventResponseReceived:
			urlMap[ev.RequestID] = ev.Response.URL
			break
		}
	})

	// Pipeline
	tasks := chromedp.Tasks{
		// 쿠팡 파트너스 홈페이지 접속
		chromedp.Navigate("https://partners.coupang.com/"),
		// 로그인 클릭
		chromedp.Click(`#app-header > div.header-toolbar > div > button:nth-child(1)`),
		// 로그인 정보 입력
		chromedp.SendKeys(`._loginIdInput`, config.Opts.CoupangId),
		chromedp.SendKeys(`._loginPasswordInput`, config.Opts.CoupangPw),
		// 로그인
		chromedp.Click(`.login__button`),
		// 상품검색 입력창이 나올때까지 대기
		chromedp.WaitVisible(`#root > div > div > div.workspace-container > div > div > div.affiliate-page > div > div > div.ant-spin-nested-loading.page-spin-container > div > div > div > div > div > div > div > div:nth-child(1) > div > div > div > div > span > input`),
		// 상품 키워드 입력
		chromedp.SendKeys(`#root > div > div > div.workspace-container > div > div > div.affiliate-page > div > div > div.ant-spin-nested-loading.page-spin-container > div > div > div > div > div > div > div > div:nth-child(1) > div > div > div > div > span > input`, config.Opts.Keyword),
		// 상품검색
		chromedp.Click(`#root > div > div > div.workspace-container > div > div > div.affiliate-page > div > div > div.ant-spin-nested-loading.page-spin-container > div > div > div > div > div > div > div > div:nth-child(1) > div > div > div > div > span > span.ant-input-suffix > button`),
		// 상품검색 대기
		chromedp.WaitReady(`#root > div > div > div.workspace-container > div > div > div.affiliate-page > div > div > div.ant-spin-nested-loading.page-spin-container > div > div > div > div > div > div > section.section-product-list > div > div.ant-spin-nested-loading > div > div > div > div`),
		fetch.Enable(),
	}
	err := chromedp.Run(ctx, &tasks)
	if err != nil {
		log.Fatal(err)
	}
	err = chromedp.Cancel(ctx)
	if err != nil {
		return
	}
}
