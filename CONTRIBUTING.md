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

## Schema notes

- `twitter` and `cfp` are optional values.
- `mode` should have one of the following values `hybrid`, `onsite`, or `virtual`

## Contributors

Although the public website does not reflect contributors yet, add yourself also to `./contributors.yaml` as shown below:

```yaml
- [GitHubUserName](your website, twitter profile or github profile)
```

For example:

```yaml
- [JohnDoe](https://john-does-tech-blog.sample)
```

## Conventional commit messages

Conventional commit messages are enforced on this repository. PRs not using conventional commit messages will not be merged. Consult [https://www.conventionalcommits.org/en/v1.0.0/](https://www.conventionalcommits.org/en/v1.0.0/) to learn how to write conventional commit messages.
