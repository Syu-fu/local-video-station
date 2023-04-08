import { type FC } from 'react';
import { Card, CardMedia, CardContent, Typography } from '@mui/material';
import { styled } from '@mui/system';
import type Video from '../types/video';

interface Props {
  video: Video;
}

const VideoListItemComponent: FC<Props> = ({ video }) => {
  return (
    <StyledCard>
      <StyledCardMedia image={video.thumbnailUrl} />
      <CardContent>
        <Typography variant="subtitle1" gutterBottom>
          {video.title}
        </Typography>
      </CardContent>
    </StyledCard>
  );
};

const StyledCard = styled(Card)({
  display: 'inline-flex',
  width: '100%',
  height: 'auto',
  margin: '8px',
  borderRadius: '8px',
  overflow: 'hidden',
  boxShadow: '0 0 10px rgba(0, 0, 0, 0.2)',
});

const StyledCardMedia = styled(CardMedia)({
  width: 120,
  height: 'auto',
});
export default VideoListItemComponent;
