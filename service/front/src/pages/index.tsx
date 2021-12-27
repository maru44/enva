import type { NextPage } from 'next'
import styles from '../styles/Home.module.css'
import {
  Box,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Typography,
} from '@mui/material'
import { Eco, Security, Share } from '@material-ui/icons'

const Home: NextPage = () => {
  return (
    <Box className={styles.container}>
      <Box mt={12}>
        <Typography textAlign="center" variant="h1" className={styles.title}>
          Welcome to <span className={styles.service}>Envassador</span>!
        </Typography>
      </Box>
      <Box mt={6}>
        <List disablePadding className={styles.list}>
          <ListItem>
            <ListItemIcon>
              <Security />
            </ListItemIcon>
            <ListItemText>
              Secure sharing of environmental variables.
            </ListItemText>
          </ListItem>
          <ListItem>
            <ListItemIcon>
              <Share />
            </ListItemIcon>
            <ListItemText>
              Share environmetal variables without annoying.
            </ListItemText>
          </ListItem>
          <ListItem>
            <ListItemIcon>
              <Eco />
            </ListItemIcon>
            <ListItemText>
              Eliminate the difference of environmental variables between
              developers in team development.
            </ListItemText>
          </ListItem>
        </List>
      </Box>
    </Box>
  )
}

export default Home
