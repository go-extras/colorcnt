# Colorcnt

Colorcnt is a utility that takes a list of image urls and processes it concurrently.
During the processing it finds top most used colors in those image and writes the results to a CSV file.

Note: this application ignores alpha channel.

## The initial task

> Bellow is a list of links leading to an image, read this list of images and find 3 most prevalent colors in the RGB scheme in hexadecimal format (#000000 - #FFFFFF) in each image, and write the result into a CSV file in a form of url,color,color,color.
>
> Please focus on speed and resources. The solution should be able to handle input files with more than a billion URLs, using limited resources (e.g. 1 CPU, 512MB RAM). Keep in mind that there is no limit on the execution time, but make sure you are utilizing the provided resources as much as possible at any time during the program execution.
>
> Answer should be posted in a git repo.

## Building

To build this application you should have go 1.13 or later installed.

In the root directory of the project run the following command to build the application:

```bash
go build
``` 

As a result you will get a binary `./colorcnt`.

## Running

Here is how you can easily run the app:

```bash
./colorcnt run --input-file=input.txt --output-file=output.csv
```

Note that you may tune additional parameters that would better correspond to your server environment:

```
$ ./colorcnt run --help
Usage:
  colorcnt [OPTIONS] run [run-OPTIONS]

Help Options:
  -h, --help                Show this help message

[run command options]
          --input-file=     Input file name (must consist of URL)
          --output-file=    Output file name
          --max-workers=    Max workers to run concurrently (default: 10)
          --max-top-colors= Max number of the top colors to choose (default: 3)
          --write-buffer=   Size of CSVWriter buffer (in records) (default: 10000)
```

Basically, you can tune the number of workers, number of colors and the size of the CSVWriter buffer.

Please make sure you set the write buffer high enough if you decide to increase the number of workers, otherwise you will hit the bottle neck of the writer.
Ideally, it's better to have multiple instances of this app writing to different locations to avoid the limitations of non-concurrent file writes. 

## Testing

Running tests:

```bash
# in the project root folder
go test ./... -race
```

## Final notes

This is just a demo application and lacks documentation, lacks some tests and of course may have bugs. Feel free letting me know through github issues ;)
