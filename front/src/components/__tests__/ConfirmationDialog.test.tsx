import { render, fireEvent, screen } from '@testing-library/react';
import ConfirmationDialogComponent from '../ConfirmationDialogComponent';

describe('ConfirmationDialog component', () => {
  test('renders the dialog with the provided title and message', () => {
    const title = 'Confirmation';
    const message = 'Are you sure you want to submit?';
    const isOpen = true;
    const onCancel = vitest.fn();
    const onConfirm = vitest.fn();

    render(
      <ConfirmationDialogComponent
        title={title}
        message={message}
        isOpen={isOpen}
        onCancel={onCancel}
        onConfirm={onConfirm}
      />
    );

    expect(screen.getByText(title)).toBeInTheDocument();
    expect(screen.getByText(message)).toBeInTheDocument();
  });

  test('calls onCancel when the cancel button is clicked', () => {
    const title = 'Confirmation';
    const message = 'Are you sure you want to submit?';
    const isOpen = true;
    const onCancel = vitest.fn();
    const onConfirm = vitest.fn();

    render(
      <ConfirmationDialogComponent
        title={title}
        message={message}
        isOpen={isOpen}
        onCancel={onCancel}
        onConfirm={onConfirm}
      />
    );

    fireEvent.click(screen.getByText('Cancel'));
    expect(onCancel).toHaveBeenCalledTimes(1);
  });

  test('calls onConfirm when the OK button is clicked', () => {
    const title = 'Confirmation';
    const message = 'Are you sure you want to submit?';
    const isOpen = true;
    const onCancel = vitest.fn();
    const onConfirm = vitest.fn();

    render(
      <ConfirmationDialogComponent
        title={title}
        message={message}
        isOpen={isOpen}
        onCancel={onCancel}
        onConfirm={onConfirm}
      />
    );

    fireEvent.click(screen.getByText('OK'));
    expect(onConfirm).toHaveBeenCalledTimes(1);
  });
});
