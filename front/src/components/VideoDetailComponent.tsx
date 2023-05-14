import { type FC } from 'react';
import EditIcon from '@mui/icons-material/Edit';
import LabelIcon from '@mui/icons-material/Label';
import { Chip, Typography, Box, IconButton } from '@mui/material';
import { Link } from 'react-router-dom';
import type Video from '../types/video';

interface Props {
  video: Video;
}

const VideoDetailComponent: FC<Props> = ({ video }) => {
  return (
    <Box sx={{ padding: '16px' }}>
      <Box
        sx={{
          display: 'flex',
          alignItems: 'center',
          marginBottom: '10px',
        }}
      >
        <Typography
          variant="h6"
          sx={{ fontWeight: 'bold', marginBottom: '10px' }}
        >
          {video.title}
        </Typography>
        <IconButton
          component={Link}
          to={`/video/edit/${video.id}`}
          sx={{ marginTop: '10px' }}
        >
          <EditIcon />
        </IconButton>
      </Box>
      {video.tags.map((tag) => {
        return (
          <Chip
            key={tag.id}
            label={tag.name}
            icon={<LabelIcon />}
            sx={{ marginRight: '5px', marginBottom: '5px' }}
          />
        );
      })}
    </Box>
  );
};

export default VideoDetailComponent;
