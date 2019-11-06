package http

import (
	"github.com/open-falcon/agent/g"
	"github.com/toolkits/file"
	"net/http"
	"path/filepath"
	"strings"
)

func configPageRoutes() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			if !file.IsExist(filepath.Join(g.Root, "/public", r.URL.Path, "index.html")) {
				http.NotFound(w, r)
				return
			}
		}
		// 启动fileServer暴露public下面的文件, 可直接访问静态文件
		http.FileServer(http.Dir(filepath.Join(g.Root, "/public"))).ServeHTTP(w, r)
	})

}
