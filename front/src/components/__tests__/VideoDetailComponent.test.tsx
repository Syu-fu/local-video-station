import { render, screen } from '@testing-library/react';
import type Video from '../../types/video';
import VideoDetailComponent from '../VideoDetailComponent';

describe('VideoDetailComponent', () => {
  const testVideo: Video = {
    id: '3C6FC606-3DC8-4EE5-AFC0-D6C20DE47AE4',
    title: 'How to Make Sushi',
    titleReading: 'How to Make Sushi',
    url: 'https://example.com/how-to-make-sushi',
    thumbnailUrl: 'https://example.com/sushi-thumbnail.jpg',
    tags: [
      {
        id: '74FF4BBA-18A5-4ABC-9164-4987021D411F',
        name: 'Sushi',
        nameReading: 'Sushi',
      },
      {
        id: 'A082A0B6-63DE-4E97-BA30-B29319863D4F',
        name: 'Cooking',
        nameReading: 'Cooking',
      },
    ],
  };

  it('renders the video title', () => {
    render(<VideoDetailComponent video={testVideo} />);
    expect(screen.getByText(testVideo.title)).toBeVisible();
  });

  it('renders the video tags with the correct icon', () => {
    render(<VideoDetailComponent video={testVideo} />);
    testVideo.tags.forEach((tag) => {
      const chip = screen.getByText(tag.name);
      expect(chip).toBeInTheDocument();
      expect(chip).toHaveClass('MuiChip-label');
    });
  });
});
