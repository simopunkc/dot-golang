package abstraction

import "dot-golang/internal/domain"

type BlogEvent interface {
	PostNews(new domain.News)
}
