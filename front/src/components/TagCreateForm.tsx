import { type FC, useState } from 'react';
import { css } from '@emotion/react';
import { Button, TextField } from '@mui/material';
import { useForm } from 'react-hook-form';
import postTagDetail from '../api/postTagDetail';
import ConfirmationDialogComponent from './ConfirmationDialogComponent';
import NotificationComponent from './NotificationComponent';

type TagFormInputs = {
  name: string;
  nameReading: string;
};

const containerStyle = css`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 2rem;
  width: 100%;
  max-width: 400px;
  box-sizing: border-box;
`;

const textFieldStyle = css`
  margin: 8px 0;
  width: 100%;
`;

const TagCreateForm: FC = () => {
  const { register, handleSubmit, reset } = useForm<TagFormInputs>({
    defaultValues: {
      name: '',
      nameReading: '',
    },
  });
  const [confirmationOpen, setConfirmationOpen] = useState(false);
  const [notificationOpen, setNotificationOpen] = useState(false);
  const [alertSeverity, setAlertSeverity] = useState<'success' | 'error'>(
    'success'
  );
  const [alertMessages, setAlertMessages] = useState<string[]>([]);

  const handleFormSubmit = () => {
    void handleSubmit(async (data: TagFormInputs) => {
      try {
        handleCloseConfirmation();
        const tag = await postTagDetail({ id: '', ...data });
        setAlertSeverity('success');
        setAlertMessages([
          `ID: ${tag.id}`,
          `Name: ${tag.name}`,
          `Name Reading: ${tag.nameReading}`,
        ]);
      } catch (error) {
        if (error instanceof Error) {
          setAlertSeverity('error');
          setAlertMessages([`An error occurred: ${error.message}`]);
        } else {
          setAlertSeverity('error');
          setAlertMessages(['An unknown error occurred.']);
        }
      } finally {
        reset();
        setNotificationOpen(true);
      }
    })();
  };

  const handleCloseNotification = () => {
    setNotificationOpen(false);
  };

  const handleOpenConfirmation = () => {
    setConfirmationOpen(true);
  };

  const handleCloseConfirmation = () => {
    setConfirmationOpen(false);
  };

  return (
    <div css={containerStyle}>
      <form onSubmit={handleSubmit(handleFormSubmit)}>
        <TextField
          label="Name"
          {...register('name', { required: true })}
          variant="outlined"
          css={textFieldStyle}
        />
        <TextField
          label="Name Reading"
          {...register('nameReading', { required: true })}
          variant="outlined"
          css={textFieldStyle}
        />
        <Button onClick={handleOpenConfirmation} variant="contained">
          Submit
        </Button>
      </form>
      <ConfirmationDialogComponent
        isOpen={confirmationOpen}
        title="Submit Confirmation"
        message="Are you sure you want to submit this form?"
        onConfirm={handleFormSubmit}
        onCancel={handleCloseConfirmation}
      />
      <NotificationComponent
        open={notificationOpen}
        severity={alertSeverity}
        messages={alertMessages}
        onClose={handleCloseNotification}
      />
    </div>
  );
};

export default TagCreateForm;
