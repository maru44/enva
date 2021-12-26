import { Delete } from '@material-ui/icons'
import {
  Card,
  Grid,
  IconButton,
  Paper,
  Tooltip,
  Typography,
} from '@mui/material'
import clsx from 'clsx'
import Link from 'next/link'
import { Project } from '../../../types/project'
import theme from '../../theme/theme'
import styles from '../../styles/project.module.css'

type props = {
  project: Project
  startDeleteFunc: () => void
}

export const ProjectListCard: React.FC<props> = ({
  project,
  startDeleteFunc,
}) => {
  return (
    <Grid item md={4} xs={6}>
      <Card
        className={clsx(styles.card, 'hrefBox')}
        component={Paper}
        variant="outlined"
      >
        <Grid container pl={2} pr={2} pt={1} pb={1}>
          <Grid
            item
            xs={12}
            display="flex"
            flexDirection="row"
            alignItems="center"
            justifyContent="space-between"
          >
            <Grid item flex={1} overflow="hidden">
              <Typography variant="h6">{project.name}</Typography>
            </Grid>
            <Grid item width={40}>
              <Tooltip title="delete project" arrow>
                <IconButton
                  className={styles.deleteIcon}
                  onClick={startDeleteFunc}
                >
                  <Delete />
                </IconButton>
              </Tooltip>
            </Grid>
          </Grid>
          <Grid item xs={12} mt={1} overflow="hidden">
            <Typography maxHeight={theme.spacing(9.5)}>
              {project.description}
            </Typography>
          </Grid>
        </Grid>
        <Link as={`/project/${project.slug}`} href={`/project/[slug]`} passHref>
          <a className="hrefBoxIn"></a>
        </Link>
      </Card>
    </Grid>
  )
}
