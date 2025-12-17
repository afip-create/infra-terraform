type WebResponse struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data,omitempty"`
    Error   *string     `json:"error,omitempty"`
}

func (r *WebResponse) ErrorOrNil() error {
    if r.Error != nil {
        return fmt.Errorf("error: %s", *r.Error)
    }
    return nil
}