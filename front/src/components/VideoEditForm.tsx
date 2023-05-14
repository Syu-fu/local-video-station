import { type FC, useState, useEffect } from 'react';
import { css } from '@emotion/react';
import {
  TextField,
  Button,
  Box,
  Input,
  InputLabel,
  FormControl,
} from '@mui/material';
import { useForm, Controller } from 'react-hook-form';
import getTagList from '../api/getTagList';
import putVideoDetail from '../api/putVideoDetail';
import type Tag from '../types/tag';
import type Video from '../types/video';
import type VideoFormInputs from '../types/videoFormInputs';
import ConfirmationDialogComponent from './ConfirmationDialogComponent';
import NotificationComponent from './NotificationComponent';
import ProgressBarComponent from './ProgressBarComponent';
import TagAutocompleteComponent from './TagAutoCompleteComponent';

const formStyle = css`
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

type VideoEditFormProps = {
  video: Video;
};

const VideoEditForm: FC<VideoEditFormProps> = ({ video }) => {
  const { control, handleSubmit, setValue, reset } = useForm<VideoFormInputs>();
  const [tags, setTags] = useState<Tag[]>([]);
  const [confirmationOpen, setConfirmationOpen] = useState(false);
  const [notificationOpen, setNotificationOpen] = useState(false);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [alertSeverity, setAlertSeverity] = useState<'success' | 'error'>(
    'success'
  );
  const [alertMessages, setAlertMessages] = useState<string[]>([]);
  const [thumbnailKey, setThumbnailKey] = useState(0);
  const [videoKey, setVideoKey] = useState(0);

  useEffect(() => {
    getTagList()
      .then((data) => {
        setTags(data);
      })
      .catch((error) => {
        console.log(error);
      });
  }, []);

  useEffect(() => {
    reset({
      id: video.id,
      title: video.title,
      titleReading: video.titleReading,
      tags: video.tags,
    });
  }, [video, reset]);

  const handleFormSubmit = () => {
    void handleSubmit(async (data: VideoFormInputs) => {
      setIsSubmitting(true);
      try {
        handleCloseConfirmation();
        const video = await putVideoDetail(data);
        setAlertSeverity('success');
        const tagNames = data.tags.map((tag) => tag.name).join(', ');
        setAlertMessages([
          `ID: ${video.id}`,
          `Title: ${video.title}`,
          `Title Reading: ${video.titleReading}`,
          `Tag: ${tagNames}`,
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
        setThumbnailKey(thumbnailKey + 1);
        setVideoKey(videoKey + 1);
        setNotificationOpen(true);
        setIsSubmitting(false);
      }
    })();
  };

  const handleThumbnailUpload = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    if (event.target.files != null && event.target.files.length > 0) {
      setValue('thumbnail', event.target.files[0]);
    }
  };

  const handleVideoUpload = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files != null && event.target.files.length > 0) {
      setValue('thumbnail', event.target.files[0]);
    }
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
    <div>
      {isSubmitting && <ProgressBarComponent />}
      <form
        css={formStyle}
        encType="multipart/form-data"
        onSubmit={handleSubmit(handleFormSubmit)}
      >
        <TextField
          label="ID"
          variant="outlined"
          css={textFieldStyle}
          disabled
          defaultValue={video.id}
        />
        <Controller
          name="title"
          control={control}
          defaultValue=""
          render={({ field }) => (
            <TextField
              label="title"
              variant="outlined"
              css={textFieldStyle}
              {...field}
            />
          )}
        />
        <Controller
          name="titleReading"
          control={control}
          defaultValue=""
          render={({ field }) => (
            <TextField
              label="titleReading"
              variant="outlined"
              css={textFieldStyle}
              {...field}
            />
          )}
        />
        <FormControl sx={{ marginTop: '16px' }}>
          <InputLabel htmlFor="thumbnail-upload" shrink>
            thumbnail
          </InputLabel>
          <Input
            key={thumbnailKey}
            id="thumbnail-upload"
            type="file"
            inputProps={{ accept: 'image/*' }}
            onChange={handleThumbnailUpload}
            sx={{ marginTop: '8px' }}
          />
        </FormControl>
        <FormControl sx={{ marginTop: '16px' }}>
          <InputLabel htmlFor="video-upload" shrink>
            video
          </InputLabel>
          <Input
            key={videoKey}
            id="video-upload"
            type="file"
            inputProps={{ accept: 'video/*' }}
            onChange={handleVideoUpload}
            sx={{ marginTop: '8px' }}
          />
        </FormControl>
        <Controller
          name="tags"
          control={control}
          defaultValue={[]}
          render={({ field }) => (
            <TagAutocompleteComponent
              tags={tags}
              selectedTags={field.value}
              onChange={field.onChange}
            />
          )}
        />
        <Box mt={2}>
          <Button
            onClick={handleOpenConfirmation}
            variant="contained"
            color="primary"
          >
            Submit
          </Button>
        </Box>
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

export default VideoEditForm;
