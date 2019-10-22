# downloads_path_provider

This Go package implements the host-side of the Flutter [downloads_path_provider](https://github.com/jld3103/downloads_path_provider) plugin.

## Usage

Import as:

```go
import downloads_path_provider "github.com/jld3103/downloads_path_provider/go"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info):

```go
flutter.AddPlugin(&downloads_path_provider.DownloadsPathProviderPlugin{}),
```
