# Awesome Tech Conferences

*Awesome Tech Conferences* is a curated list of tech conferences.

## Contributing

Feel free to submit new tech conferences by adding them to `./confs.yaml`. For every new conference provide the following metadata:

```yaml
name: Conference Name
description: Description about the conference
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

### Schema notes

- `twitter` and `cfp` are optional values.
- `mode` should have one of the following values `hybrid`, `onsite`, or `virtual`
- `description` is truncated after `285` characters

### Contributors

Although the public website does not reflect contributors yet, add yourself also to `./contributors.md` as shown below:

```md
- [GitHubUserName](your website, twitter profile or github profile)
```

For example:

```md
- [JohnDoe](https://john-does-tech-blog.sample)
```
