import { Badge, Typography } from '@mui/material'
import { InvitationStatus } from '../../../../types/org'
import styles from '../../../styles/org.module.css'

type props = {
  status: InvitationStatus
}

export const InvitationStatusBadge: React.FC<props> = ({ status }) => {
  return (
    <span
      className={`${styles.invitationStatusLabel} ${
        styles[`invitation-${status}`]
      }`}
    >
      {status}
    </span>
  )
}
