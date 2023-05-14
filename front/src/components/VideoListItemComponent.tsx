import { type FC } from 'react';
import { Card, CardMedia, CardContent, Typography } from '@mui/material';
import { styled } from '@mui/system';
import { Link } from 'react-router-dom';
import type Video from '../types/video';

interface Props {
  video: Video;
}

const VideoListItemComponent: FC<Props> = ({ video }) => {
  return (
    <Link to={`/video/${video.id}`}>
      <StyledCard>
        <StyledCardMedia image={video.thumbnailUrl} />
        <CardContent>
          <Typography variant="subtitle1" gutterBottom>
            {video.title}
          </Typography>
        </CardContent>
      </StyledCard>
    </Link>
  );
};

const StyledCard = styled(Card)(() => ({
  display: 'flex',
  flexDirection: 'column',
  width: '100%',
  height: '100%',
  margin: '8px',
  borderRadius: '8px',
  overflow: 'hidden',
  boxShadow: '0 0 10px rgba(0, 0, 0, 0.2)',
  flexShrink: 0,
  flexGrow: 1,
}));

const StyledCardMedia = styled(CardMedia)(() => ({
  width: '100%',
  height: 0,
  paddingTop: '56.25%',
  flexShrink: 0,
}));

export default VideoListItemComponent;
