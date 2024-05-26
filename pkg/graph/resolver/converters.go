package resolver

import (
	"github.com/nareshbhatia/movie-magic-gateway-go/pkg/graph/model"
	movie "github.com/nareshbhatia/movie-magic-services-go/gen/go/models/movie/v1"
	movieService "github.com/nareshbhatia/movie-magic-services-go/gen/go/movie/v1"
)

// ----- fromModelMoviesRequest -----

func fromModelMoviesRequest(input *model.MoviesRequest) *movieService.ListMoviesRequest {
	var sort movieService.SortParam
	if input.Sort != nil {
		sort = fromModelSortParam(*input.Sort)
	}

	return &movieService.ListMoviesRequest{
		Filters:  fromModelFilterParams(input.Filters),
		Sort:     sort,
		PageSpec: fromModelPageSpec(input.PageSpec),
	}
}

func fromModelFilterParams(filters *model.FilterParams) *movieService.FilterParams {
	if filters == nil {
		return nil
	}

	certs := make([]movie.Certificate, len(filters.Certs))
	for i, cert := range filters.Certs {
		certs[i] = fromModelCertificate(cert)
	}

	return &movieService.FilterParams{
		Certs: certs,
	}
}

func fromModelSortParam(sort model.SortParam) movieService.SortParam {
	switch sort {
	case model.SortParamSortParamRankAsc:
		return movieService.SortParam_SORT_PARAM_RANK_ASC
	default:
		return movieService.SortParam_SORT_PARAM_RANK_ASC
	}
}

func fromModelPageSpec(pageSpec *model.PageSpec) *movieService.PageSpec {
	if pageSpec == nil {
		return nil
	}

	return &movieService.PageSpec{
		Page:    int32(pageSpec.Page),
		PerPage: int32(pageSpec.PerPage),
	}
}

// ----- toModelMovie -----

func toModelMovie(movie *movie.Movie) *model.Movie {
	return &model.Movie{
		ID:          movie.Id,
		Name:        movie.Name,
		Cast:        toModelCastMembers(movie.Cast),
		Certificate: toModelCertificate(movie.Certificate),
		Genres:      movie.Genres,
		Rank:        int(movie.Rank),
		RatingsSummary: &model.MovieRatingsSummary{
			AggregateRating: movie.RatingsSummary.AggregateRating,
			VoteCount:       int(movie.RatingsSummary.VoteCount),
		},
		ReleaseYear: int(movie.ReleaseYear),
	}
}

func toModelCastMembers(cast []*movie.CastMember) []*model.CastMember {
	castMembers := make([]*model.CastMember, len(cast))
	for i, member := range cast {
		castMembers[i] = &model.CastMember{
			Person: &model.Person{
				ID:   member.PersonId,
				Name: "Actor " + member.PersonId,
			},
			Characters: member.Characters,
		}
	}
	return castMembers
}

// ----- to/from ModelCertificate -----

func fromModelCertificate(certificate model.Certificate) movie.Certificate {
	switch certificate {
	case model.CertificateCertificateUnspecified:
		return movie.Certificate_CERTIFICATE_UNSPECIFIED
	case model.CertificateCertificateG:
		return movie.Certificate_CERTIFICATE_G
	case model.CertificateCertificateNr:
		return movie.Certificate_CERTIFICATE_NR
	case model.CertificateCertificatePg13:
		return movie.Certificate_CERTIFICATE_PG_13
	case model.CertificateCertificatePg:
		return movie.Certificate_CERTIFICATE_PG
	case model.CertificateCertificateR:
		return movie.Certificate_CERTIFICATE_R
	default:
		return movie.Certificate_CERTIFICATE_UNSPECIFIED
	}
}

func toModelCertificate(certificate movie.Certificate) model.Certificate {
	switch certificate {
	case movie.Certificate_CERTIFICATE_UNSPECIFIED:
		return model.CertificateCertificateUnspecified
	case movie.Certificate_CERTIFICATE_G:
		return model.CertificateCertificateG
	case movie.Certificate_CERTIFICATE_NR:
		return model.CertificateCertificateNr
	case movie.Certificate_CERTIFICATE_PG_13:
		return model.CertificateCertificatePg13
	case movie.Certificate_CERTIFICATE_PG:
		return model.CertificateCertificatePg
	case movie.Certificate_CERTIFICATE_R:
		return model.CertificateCertificateR
	default:
		return model.CertificateCertificateUnspecified
	}
}

// ----- toModelPageInfo -----

func toModelPageInfo(pageInfo *movieService.PaginationInfo) *model.PaginationInfo {
	return &model.PaginationInfo{
		TotalPages:      int(pageInfo.TotalPages),
		TotalItems:      int(pageInfo.TotalItems),
		Page:            int(pageInfo.Page),
		PerPage:         int(pageInfo.PerPage),
		HasNextPage:     pageInfo.HasNextPage,
		HasPreviousPage: pageInfo.HasPreviousPage,
	}
}
