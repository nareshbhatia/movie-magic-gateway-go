package resolver

import (
	"github.com/nareshbhatia/movie-magic-gateway-go/pkg/graph/model"
	moviev1 "github.com/nareshbhatia/movie-magic-services-go/gen/go/models/movie/v1"
)

func convertCertificate(certificate moviev1.Certificate) model.Certificate {
	switch certificate {
	case moviev1.Certificate_CERTIFICATE_UNSPECIFIED:
		return model.CertificateCertificateUnspecified
	case moviev1.Certificate_CERTIFICATE_G:
		return model.CertificateCertificateG
	case moviev1.Certificate_CERTIFICATE_NR:
		return model.CertificateCertificateNr
	case moviev1.Certificate_CERTIFICATE_PG_13:
		return model.CertificateCertificatePg13
	case moviev1.Certificate_CERTIFICATE_PG:
		return model.CertificateCertificatePg
	case moviev1.Certificate_CERTIFICATE_R:
		return model.CertificateCertificateR
	default:
		return model.CertificateCertificateUnspecified
	}
}

func convertCastMembers(cast []*moviev1.CastMember) []*model.CastMember {
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
