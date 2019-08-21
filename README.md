# Random Stylesheets

### Highlights

- produces random stylesheet for randomly ordered html
- functions and endpoints accept a `seed` argument for reproducible output
- includes a server that outputs html+css
- includes a command to collect a series of screenshots and structured data (as json) of the generated stylesheet

### Goals/TODO

- have generators for all css properties (currently only a very small set)
- include enough data about the generated css to enable machine learning
- run everything in docker
- fully reproducible output (must be possible to go back in time and generate a bigger data set)
- drop ready-made datasets somewhere in a bucket
