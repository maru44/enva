import { Apartment, ArrowBack, History, Mail } from '@material-ui/icons'
import { Box, Grid, Icon, IconButton, Tooltip, Typography } from '@mui/material'
import { NextPage } from 'next'
import { useRouter } from 'next/router'
import { useState } from 'react'
import useSWR from 'swr'
import { useRequireLogin } from '../../../hooks/useRequireLogin'
import { OrgResponseBody, OrgsResponseBody } from '../../../http/body/org'
import { projectsResponseBody } from '../../../http/body/project'
import { fetcherGetFromApiUrl, GetPath } from '../../../http/fetcher'
import { PageProps } from '../../../types/page'
import { Project } from '../../../types/project'
import { UserUserTypes } from '../../../types/user'
import { CommonListCard } from '../../components/CommonListCard'
import { DeleteModal } from '../../components/DeleteModal'
import { InvitationHistoryModal } from '../../components/form/org/InvitationHistoryModal'
import { InviteFormModal } from '../../components/form/org/InviteFormModal'
import { MembersList } from '../../components/form/org/MembersList'
import styles from '../../styles/project.module.css'

const OrgDetail: NextPage<PageProps> = (props) => {
  useRequireLogin()

  const router = useRouter()
  const slug = router.query.slug
  const [inviteFormOpen, setInviteFormOpen] = useState<boolean>(false)
  const [isHistoryOpen, setIsHistoryOpen] = useState<boolean>(false)

  const { data, error } = useSWR<OrgResponseBody, ErrorConstructor>(
    `${GetPath.ORG_DETAIL}?slug=${slug}`,
    fetcherGetFromApiUrl
  )

  if (error) {
    router.push('/500')
    return <div></div>
  }

  if (data?.data) {
    const org = data.data.org
    const userType = data.data.current_user_type!

    return (
      <Box>
        <Box mt={6}>
          <Box
            display="flex"
            flexDirection="row"
            alignItems="center"
            justifyContent="space-between"
          >
            <Box display="flex" flexDirection="row" alignItems="center">
              <Box mr={2}>
                <IconButton
                  onClick={() => {
                    router.back()
                  }}
                >
                  <ArrowBack />
                </IconButton>
              </Box>
              <Box display="flex" flexDirection="row" alignItems="center">
                <Icon>
                  <Apartment />
                </Icon>
                <Typography variant="h5">{org!.name}</Typography>
              </Box>
            </Box>
          </Box>
          <Box mt={2} pl={1} pr={1}>
            {org.description && (
              <Box>
                <Typography>{org!.description}</Typography>
              </Box>
            )}
            <Box mt={6}>
              <OrgProjects id={org!.id} slug={org!.slug} />
            </Box>
            <Box mt={4}>
              <Typography variant="h6">{org.user_count} Members</Typography>
            </Box>
            <MembersList id={org.id} currentUserType={userType}></MembersList>
            {UserUserTypes.includes(userType) && (
              <Box mt={4}>
                <Box>
                  <Typography variant="h6">Invitations</Typography>
                </Box>
                <Box mt={2} display="flex" flexDirection="row">
                  <Box mr={2}>
                    <Tooltip title="invite" arrow placement="bottom">
                      <IconButton onClick={() => setInviteFormOpen(true)}>
                        <Mail />
                      </IconButton>
                    </Tooltip>
                  </Box>
                  <Box>
                    <Tooltip title="history" arrow placement="bottom">
                      <IconButton onClick={() => setIsHistoryOpen(true)}>
                        <History />
                      </IconButton>
                    </Tooltip>
                  </Box>
                </Box>
              </Box>
            )}
          </Box>
        </Box>
        <InviteFormModal
          orgId={org!.id}
          orgName={org!.name}
          isOpen={inviteFormOpen}
          onClose={() => {
            setInviteFormOpen(false)
          }}
        />
        <InvitationHistoryModal
          orgId={org!.id}
          isOpen={isHistoryOpen}
          onClose={() => {
            setIsHistoryOpen(false)
          }}
        />
      </Box>
    )
  }

  return <div></div>
}

type orgProjectsProps = {
  id: string
  slug: string
}

const OrgProjects: React.FC<orgProjectsProps> = ({ id, slug }) => {
  const { data, error } = useSWR<projectsResponseBody, ErrorConstructor>(
    `${GetPath.PROJECT_LIST_ORG}?id=${id}`,
    fetcherGetFromApiUrl
  )
  const [project, setProject] = useState<Project | undefined>(undefined)
  const [isOpen, setIsOpen] = useState<boolean>(false)

  if (error) return <div></div>

  return (
    <Grid container rowSpacing={2} columnSpacing={2}>
      {data &&
        data.data &&
        data.data.map((p, i) => (
          <CommonListCard
            info={p}
            key={i}
            linkAs={`/project/${slug}/${p.slug}`}
            linkHref="/project/[...slug]"
            styles={styles}
            startDeleteFunc={() => {
              setProject(p)
              setIsOpen(true)
            }}
          />
        ))}
      <DeleteModal
        url={`${GetPath.PROJECT_DELETE}?projectId=${project?.id}`}
        isOpen={isOpen}
        mutateKey={`${GetPath.PROJECT_LIST_ORG}?id=${id}`}
        Message={<Typography variant="h5">Delete {project?.name}?</Typography>}
        onClose={() => setIsOpen(false)}
      />
    </Grid>
  )
}

export default OrgDetail
