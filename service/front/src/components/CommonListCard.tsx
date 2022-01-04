import { Delete } from '@material-ui/icons'
import {
  Box,
  Card,
  Grid,
  IconButton,
  Paper,
  Tooltip,
  Typography,
} from '@mui/material'
import clsx from 'clsx'
import Link from 'next/link'
import { ReactNode } from 'react'
import { BaseInformation } from '../../types/information'
import theme from '../theme/theme'

type props = {
  info: BaseInformation
  linkAs: string
  linkHref: string
  styles: { readonly [key: string]: string }
  icon?: ReactNode
  startDeleteFunc?: () => void
}

export const CommonListCard: React.FC<props> = ({
  info,
  linkAs,
  linkHref,
  styles,
  icon,
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
            {icon && (
              <Box width={40}>
                <IconButton>{icon}</IconButton>
              </Box>
            )}
            <Box flex={1} overflow="hidden">
              <Typography variant="h6">{info.name}</Typography>
            </Box>
            <Box width={40}>
              <Tooltip title="delete project" arrow>
                <IconButton
                  className={styles.deleteIcon}
                  onClick={startDeleteFunc}
                >
                  <Delete />
                </IconButton>
              </Tooltip>
            </Box>
          </Grid>
          <Grid item xs={12} mt={1} overflow="hidden">
            <Typography maxHeight={theme.spacing(9.5)}>
              {info.description}
            </Typography>
          </Grid>
        </Grid>
        <Link as={linkAs} href={linkHref} passHref>
          <a className="hrefBoxIn"></a>
        </Link>
      </Card>
    </Grid>
  )
}
