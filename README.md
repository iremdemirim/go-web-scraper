# Go Web Scraper 🕷️

Bu araç, terminalden girilen bir web sitesini ziyaret eder ve verileri (sitenin ekran görüntüsü, HTML kodu, linkler) otomatik olarak kaydeder.

## Özellikler

1.  **Ekran Görüntüsü:** Sayfanın tam boy fotoğrafını çeker (`ekrangor.png`).
2.  **Kaynak Kod:** Sitenin HTML kodlarını indirir (`site.html`).
3.  **Linkler:** Sayfadaki tüm bağlantıları listeler (`links.txt`).

## Gereksinimler

* **Go** (Golang)
* **Google Chrome** Tarayıcısı
* **Linux** İşletim Sistemi

## Kurulum

Proje klasöründe terminali açın ve gerekli paketleri indirin:

```bash
go mod tidy

## Kullanım

go run main.go https://site.com
