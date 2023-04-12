/**
 * Created by YuYoung on 2023/4/12
 * Description: 链接 ORM Model
 */

package mysql

type Link struct {
	ShortLink string `json:"shortLink" gorm:"primaryKey"`
	LongLink  string `json:"longLink"`
}
