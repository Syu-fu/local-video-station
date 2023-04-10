import { type FC } from 'react';
import LabelIcon from '@mui/icons-material/Label';
import { Chip, Typography } from '@mui/material';
import type Video from '../types/video';

interface Props {
  video: Video;
}

const VideoDetailComponent: FC<Props> = ({ video }) => {
  return (
    <>
      <Typography
        variant="h6"
        sx={{ fontWeight: 'bold', marginBottom: '10px' }}
      >
        {video.title}
      </Typography>
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
    </>
  );
};

export default VideoDetailComponent;
