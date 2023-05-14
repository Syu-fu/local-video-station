import { type FC } from 'react';
import { Grid } from '@mui/material';
import type Video from '../types/video';
import VideoListItemComponent from './VideoListItemComponent';

interface Props {
  videos: Video[];
}

const VideoListComponent: FC<Props> = ({ videos }) => {
  return (
    <Grid container spacing={2}>
      {videos.map((video) => (
        <Grid item xs={12} sm={6} key={video.id}>
          <VideoListItemComponent video={video} />
        </Grid>
      ))}
    </Grid>
  );
};

export default VideoListComponent;
