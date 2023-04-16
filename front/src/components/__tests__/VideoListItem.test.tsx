import { render, screen } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
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
    tags: [],
  };

  it('should render video thumbnail, title, and url', () => {
    render(
      <MemoryRouter>
        <VideoListItemComponent video={video} />
      </MemoryRouter>
    );
    const cardMediaElement = screen.getByRole('img');
    const title = screen.getByText(video.title);

    expect(title).toHaveTextContent(video.title);
    expect(cardMediaElement).toHaveStyle(
      `background-image: url(${testThumbnailUrl});`
    );
  });
});

test('renders video title and link', () => {
  const video: Video = {
    id: '3C6FC606-3DC8-4EE5-AFC0-D6C20DE47AE4',
    title: 'How to Make Sushi',
    titleReading: 'How to Make Sushi',
    url: 'http://test.com/movie.mp4',
    thumbnailUrl: 'http://test.com/thumbnail.jpg',
    tags: [],
  };

  render(
    <MemoryRouter>
      <VideoListItemComponent video={video} />
    </MemoryRouter>
  );

  const linkElement = screen.getByRole('link');
  expect(linkElement).toHaveAttribute(
    'href',
    '/video/3C6FC606-3DC8-4EE5-AFC0-D6C20DE47AE4'
  );

  const titleElement = screen.getByText('How to Make Sushi');
  expect(titleElement).toBeInTheDocument();
});
