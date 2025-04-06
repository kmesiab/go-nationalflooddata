# National Flood Data API Wrapper

[National Flood Data Official Website](https://docs.nationalflooddata.com/)

![Golang](https://img.shields.io/badge/Go-00add8.svg?labelColor=171e21&style=for-the-badge&logo=go)

![Build](https://github.com/kmesiab/go-nationalflooddata/actions/workflows/go-build.yml/badge.svg)
![Build](https://github.com/kmesiab/go-nationalflooddata/actions/workflows/go-lint.yml/badge.svg)
![Build](https://github.com/kmesiab/go-nationalflooddata/actions/workflows/go-test.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/kmesiab/go-nationalflooddatar)](https://goreportcard.com/report/github.com/kmesiab/go-nationalflooddata)

This Go package provides a client for interacting with the National Flood Data
API, specifically designed to facilitate access to flood-related data for
research, planning, and emergency response. The package includes methods
for querying flood data, retrieving flood maps, and processing batch requests.

## Features

- Query flood data for specific locations using addresses or coordinates.
- Retrieve raw flood map polygons in GeoJSON format.
- Process batch requests for multiple flood data queries.
- Sanitize API responses to handle inconsistencies and access restrictions.

## Installation

To install the package, use the following command:

```bash
go get -u github.com/kmesiab/go-nationalflooddata
```

## Usage

### Initializing the Service

To start using the API, you need to create a new service client with your API
key:

```go
import (
    nfd "github.com/kmesiab/go-nationalflooddata"
)

func main() {
    svc := nfd.NewService("your-api-key")
    // Use the service client to make API requests
}
```

### Querying Flood Data

You can query flood data for a specific location using the `GetFloodData`
method. This method requires a context and a set of options specifying the
search parameters.

```go
ctx := context.Background()
opts := nfd.FloodDataOptions{
    SearchType: nfd.SearchTypeAddressParcel,
    Address:    "430 Australian Ave Palm Beach, FL 33480",
    Elevation:  true,
    LOMA:       true,
    Property:   true,
    Parcel:     true,
}

floodData, err := svc.GetFloodData(ctx, opts)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("FEMA flood zone: %+v\n", floodData.Result.FloodFldHazAr)
```

### Retrieving Flood Map Raw Data

To retrieve raw flood map polygons, use the `GetFloodMapRaw` method. This
method returns large GeoJSON content representing the flood map.

```go
opts := nfd.FloodMapRawOptions{
    Lat:       34.071783,
    Lng:       -118.2596,
    Size:      0.08,
    GeoJSON:   true,
    ExcludeX:  false,
    Elevation: true,
}

floodMapContent, err := svc.GetFloodMapRaw(ctx, opts)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Flood Map Content: %+v\n", floodMapContent)
```

### Processing Batch Requests

The `GetFloodDataBatch` method allows you to process multiple flood data
queries in a single batch request. It returns a `FloodDataBatch` containing
a batch ID and a URL to poll for results.

```go
batchReq := nfd.BatchDataRequest{
    Requests: []nfd.BatchRequest{
        {
            ID:         "req1",
            SearchType: nfd.SearchTypeAddressParcel,
            Address:    "430 Australian Ave Palm Beach FL 33480",
            Elevation:  true,
        },
        {
            ID:         "req2",
            SearchType: nfd.SearchTypeCoord,
            Lat:        "34.071783",
            Lng:        "-118.2596",
            Elevation:  false,
        },
    },
}

batchResp, err := svc.GetFloodDataBatch(ctx, batchReq)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Batch ID: %s - poll results at: %s\n", batchResp.BatchID, batchResp.Result)
```

## Retrieving Static Flood Map

To retrieve a static flood map image, use the `GetStaticFloodMap` method.
This method returns the image data as a byte slice.

```go
staticMapOpts := nfd.StaticMapOptions{
    Lat:        34.071783,
    Lng:        -118.2596,
    Height:     600,
    Width:      800,
    ShowMarker: true,
    ShowLegend: true,
    Zoom:       13,
}

staticMap, err := svc.GetStaticFloodMap(ctx, staticMapOpts)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Static Map Image Data: %d bytes\n", len(staticMap))
```

## Retrieving Dynamic Flood Map

To retrieve a dynamic flood map, use the `GetDynamicFloodMap` method. This
method returns the HTML content as a string.

```go
dynamicMap, err := svc.GetDynamicFloodMap(ctx, 
	"your-api-key", 
	34.071783, 
	-118.2596, 
	13, 
	true,
)

if err != nil {
    log.Fatal(err)
}
fmt.Printf("Dynamic Map HTML: %s\n", dynamicMap)
```

## Error Handling

The package provides custom error types for handling different API error
responses, such as

- `InvalidRequestError`
- `AuthenticationError`
- `NoDataAvailableError`
- `LocationNotFoundError`
- `ParcelNotFoundError`
- `InternalServerError`

## 🧼 Response Sanitization

The package includes a `sanitizeResponse` function to clean up
API responses.

This is necessary because the API may return responses with trailing spaces or
access restrictions. For instance, if your API key lacks certain privileges,
the API might return "Access Denied" strings instead of excluding the data. The
`sanitizeResponse` function trims these spaces and replaces "Access Denied" with
`nil`, ensuring that the data is clean and consistent for further processing.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE)
file for details.

## Documentation

For more detailed information about the National Flood Data API, refer to
the [official API documentation](https://docs.nationalflooddata.com/dataservice/v3/index.html#operation/getFloodData).
