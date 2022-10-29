# Contributing

Feel free to submit new tech conferences by adding them to `./confs-{{YEAR}}.yaml`. For every new conference provide the following metadata:

```yaml
name: Conference Name
start: 2022-12-31
end: 2023-01-02
url: https://conference-website.org
twitter: ConferenceTwitterHandle
cfp:
  url: https://conference-cfp.com
  start: 2022-10-19
  end: 2022-12-10
city: Berlin
country: 🇩🇪
mode: hybrid
topics: 
  - Azure
```
## DevContainer

This repo also contains a DevContainer specification. The provided DevContainer contains all languages, tools, and fundamental VSCode Extensions required to build and run the "website" on your machine.

## Running the website locally

To run the "website" locally, [Hugo](https://gohugo.io) and Node.JS (for building the CSS) is required. Alternatively, use the provided DevContainer. It contains all necessary dependencies to buld and run the "website" locally without installing anything on your local machine (expecting Docker and VSCode to be installed).

### Building CSS

The necessary CSS for the website can be built from the `themes/conferences` subfolder as shown here:

```bash
cd themes/conferences
npm i
npm run build
```

There is an additional npm task (`npm run watch`) that could be used to build CSS continuously.

### Running the Website

Although `hugo` CLI is fairly simple, the website requires all Conference data being merged together as `confs.yaml` in the `data` folder. The `Makefile` in the root folder provides a simple `serve` task that merges the data files and starts the website using `hugo`.

```bash
# run the website
make serve
```

## Schema notes

- `twitter` and `cfp` are optional values.
- `mode` should have one of the following values `hybrid`, `onsite`, or `virtual`

## Contributors

Although the public website does not reflect contributors yet, add yourself also to `./contributors.yaml` as shown below:

```markdown
- GitHubUserName
```

For example:

```markdown
- ThorstenHans
```

## Conventional commit messages

Conventional commit messages are enforced on this repository. PRs not using conventional commit messages will not be merged. Consult [https://www.conventionalcommits.org/en/v1.0.0/](https://www.conventionalcommits.org/en/v1.0.0/) to learn how to write conventional commit messages.
