import { type FC } from 'react';
import { Box } from '@mui/material';
import type Video from '../types/video';
import VideoListItemComponent from './VideoListItemComponent';

interface Props {
  videos: Video[];
}

const VideoListComponent: FC<Props> = ({ videos }) => {
  return (
    <Box
      sx={{
        display: 'flex',
        flexWrap: 'wrap',
        boxSizing: 'border-box',
        width: '100%',

        paddingRight: '8px',
        overflowX: 'hidden',
      }}
    >
      {videos.map((video) => (
        <Box
          key={video.id}
          sx={{
            flexGrow: 1,
            flexBasis: { xs: '100%', sm: 'calc(50% - 8px)' },
            minWidth: { sm: '300px' },
            boxSizing: 'border-box',
            paddingLeft: '4px',
            paddingRight: '4px',
            marginBottom: '16px',
          }}
        >
          <VideoListItemComponent video={video} />
        </Box>
      ))}
    </Box>
  );
};

export default VideoListComponent;
