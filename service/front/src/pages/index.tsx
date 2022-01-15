import type { NextPage } from 'next'
import styles from '../styles/Home.module.css'
import {
  Box,
  Card,
  Icon,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Typography,
} from '@mui/material'
import {
  Apartment,
  Eco,
  Group,
  Lock,
  MonetizationOn,
  NetworkWifi,
  Pages,
  Security,
  Share,
  Timer,
} from '@material-ui/icons'
import { ReactNode } from 'react'

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
      <Box mt={12}>
        <Typography variant="h4">
          <b>Demo</b>
        </Typography>
      </Box>
      <Box mt={12}>
        <Box>
          <Typography variant="h4" textAlign="center">
            <b>Payments</b>
          </Typography>
        </Box>
        <Box>
          <Box mt={4}>
            <Typography variant="h5" textAlign="center">
              <b>For User</b>
            </Typography>
          </Box>
          <Box mt={2} display="flex" justifyContent="center">
            <PaymentCard
              planName="Free"
              messages={[
                {
                  icon: <Apartment />,
                  message: '1 Org',
                },
                {
                  icon: <Pages />,
                  message: '5 Projects',
                },
              ]}
            />
            <PaymentCard
              planName="Paid"
              messages={[{ icon: <Timer />, message: 'Coming soon' }]}
            ></PaymentCard>
          </Box>
          <Box mt={4}>
            <Typography variant="h5" textAlign="center">
              <b>For Organization</b>
            </Typography>
          </Box>
          <Box mt={2} display="flex">
            <PaymentCard
              planName="Free"
              messages={[
                { icon: <Group />, message: '5 Members' },
                { icon: <Pages />, message: '5 Projects' },
              ]}
            />
            <PaymentCard
              planName="Paid"
              messages={[{ icon: <Timer />, message: 'Coming soon' }]}
            ></PaymentCard>
          </Box>
        </Box>
      </Box>
      <Box mt={12}>
        <Box>
          <Typography variant="h4" textAlign="center">
            <b>Planning</b>
          </Typography>
        </Box>
        <Box mt={6}>
          <List disablePadding className={styles.list}>
            <ListItem>
              <ListItemIcon>
                <MonetizationOn />
              </ListItemIcon>
              <ListItemText>Paid plan.</ListItemText>
            </ListItem>
            <ListItem>
              <ListItemIcon>
                <NetworkWifi />
              </ListItemIcon>
              <ListItemText>
                Relay with AWS KMS (Key Management Service) and GCP Secret
                Manager.
              </ListItemText>
            </ListItem>
            <ListItem>
              <ListItemIcon>
                <Lock />
              </ListItemIcon>
              <ListItemText>
                Public-Key authentication for cli (now only with username +
                Cli-Password authentication).
              </ListItemText>
            </ListItem>
          </List>
        </Box>
      </Box>
    </Box>
  )
}

type iconMessage = {
  icon?: ReactNode
  message: string
}

type cardProps = {
  planName: string
  messages: iconMessage[]
}

const PaymentCard: React.FC<cardProps> = ({ planName, messages }) => {
  return (
    <Box m={1}>
      <Card variant="outlined">
        <Box p={2} width={200} height={200}>
          <Box>
            <Typography variant="h5" textAlign="center">
              {planName}
            </Typography>
          </Box>
          <Box mt={4}>
            {messages &&
              messages.map((m, i) => (
                <Box mt={2} key={i} display="flex" alignItems="center">
                  {m.icon && <Icon>{m.icon}</Icon>}
                  <Typography ml={2} variant="h6">
                    {m.message}
                  </Typography>
                </Box>
              ))}
          </Box>
        </Box>
      </Card>
    </Box>
  )
}

export default Home
