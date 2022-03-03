package uri


import (
  "service/iana"

  "strings"
  "fmt"
  "regexp"

  "errors"
)


type URI struct {
  scheme, userinfo, host, port, authority, path, pathquery, fullpath, query, fragment string;
}

func (this URI) Scheme() string {
  return this.scheme;
}

func (this URI) Userinfo() string {
  return this.userinfo;
}

func (this URI) Host() string {
  return this.host;
}

func (this URI) Port() string {
  return this.port;
}

func (this URI) Authority() string {
  return this.authority
}

func (this URI) Path() string {
  return this.path;
}

func (this URI) PathQuery() string {
  return this.pathquery;
}

func (this URI) FullPath() string {
  return this.fullpath
}

func (this URI) Query() string {
  return this.query;
}

func (this URI) Fragment() string {
  return this.fragment
}


func NewURI(resource string) (*URI, error) {
  regex := regexp.MustCompile("([a-zA-Z0-9]+)://([^@]+@)?([a-zA-Z0-9._-]+)(:[0-9]+)?(/[^\\?]+)?(\\?[^\\#]+)?(\\#.*)?");
  matches := regex.FindAllStringSubmatch(resource, -1);
  if len(matches) == 0 {
    return nil, errors.New("uri.NewURI(string): the URL passed is invalid")
  }

  scheme := matches[0][1]

  userinfo := strings.Replace(matches[0][2], "@", "", 1)

  host := matches[0][3]

  port := func() string {
    if matches[0][4] = strings.Replace(matches[0][4], ":", "", 1); matches[0][4] == "" {
      if service := iana.ServiceByName(scheme); service != nil {
        return fmt.Sprintf("%d", service.Port())
      }

      return ""
    }

    return matches[0][4]
  }()

  authority := func() string {
    return fmt.Sprintf("%s%s%s%s%s",
                 userinfo, func() string { if userinfo != "" { return "@" }; return ""; }(),
                 host,
                 func() string { if port != "" { return ":" }; return ""; }(), port)
  }()

  path := func() string {
    if matches[0][5] == "" { return "/" };
  
    return matches[0][5]
  }()

  pathquery := func() string {
    if matches[0][6] = strings.Replace(matches[0][6], "?", "", 1); (/* path != "/" && */ matches[0][6] != "") {
      return fmt.Sprintf("%s?%s", path, matches[0][6])
    }

    return path
  }()

  fullpath := func() string {
    if matches[0][7] = strings.Replace(matches[0][7], "#", "", 1); matches[0][7] != "" {
      return fmt.Sprintf("%s#%s", pathquery, matches[0][7])
    }

    return pathquery
  }()

  query := matches[0][6]

  fragment := matches[0][7]

  return &URI{scheme, userinfo, host, port, authority, path, pathquery, fullpath, query, fragment}, nil;
}
