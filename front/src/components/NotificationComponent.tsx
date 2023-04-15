import { type FC } from 'react';
import { Snackbar, Alert } from '@mui/material';

type NotificationProps = {
  open: boolean;
  severity: 'success' | 'error';
  messages: string[];
  onClose: () => void;
};

const Notification: FC<NotificationProps> = ({
  open,
  severity,
  messages,
  onClose,
}) => {
  return (
    <Snackbar open={open} autoHideDuration={6000} onClose={onClose}>
      <Alert onClose={onClose} severity={severity} sx={{ width: '100%' }}>
        {messages.map((message, index) => (
          <p key={index}>{message}</p>
        ))}
      </Alert>
    </Snackbar>
  );
};

export default Notification;
