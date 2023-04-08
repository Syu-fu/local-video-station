import { render, screen } from '@testing-library/react';
import type Video from '../../types/video';
import VideoListItemComponent from '../VideoListItemComponent';

describe('VideoListItem', () => {
  const testThumbnailUrl = 'http://test.com/thumbnail.jpg';
  const video: Video = {
    id: '12345-67890-12345-67890',
    title: 'How to Make Sushi',
    titleReading: 'How to Make Sushi',
    url: 'http://test.com/movie.mp4',
    thumbnailUrl: testThumbnailUrl,
  };

  it('should render video thumbnail, title, and url', () => {
    render(<VideoListItemComponent video={video} />);
    const cardMediaElement = screen.getByRole('img');
    const title = screen.getByText(video.title);

    expect(title).toHaveTextContent(video.title);
    expect(cardMediaElement).toHaveStyle(
      `background-image: url(${testThumbnailUrl});`
    );
  });
});
