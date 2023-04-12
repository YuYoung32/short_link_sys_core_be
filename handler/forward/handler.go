/**
 * Created by YuYoung on 2023/4/12
 * Description: 转发请求
 */

package forward

import (
	"net/http"
)

func ForwardHandler(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.URL.Path[1:]
	http.Redirect(w, r, mapping(shortUrl), http.StatusMovedPermanently)
}
