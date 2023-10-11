.PHONY: debug
run:
	go run main.go fluid.go types.go lerp.go

.PHONY: run
run:
	mkdir -p images
	go build -o build/fluid main.go fluid.go types.go lerp.go
	./build/fluid

.PHONY: remove-images
remove-images:
	rm -rf images

.PHONY: create-mp4
create-mp4:
	ffmpeg -framerate 60 -i images/%03d.png -c:v libx264 -profile:v high -crf 20 -pix_fmt yuv420p output.mp4