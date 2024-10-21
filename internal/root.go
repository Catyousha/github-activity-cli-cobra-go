package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tenessine/github-activity/internal/domain"
)

func GetGithubActivity(username string) (error) {
	r, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events", username))
	if err != nil {
		return nil
	}
	defer r.Body.Close()

	d, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	acts := []domain.Activity{}
	err = json.Unmarshal(d, &acts)
	if err != nil {
		return err
	}
	fmt.Println("Output:")
	for _, act := range acts {
		switch act.Type {
		case "CommitCommentEvent":
			fmt.Printf("- Comment on commit %s\n", act.Repo.Name)
		case "CreateEvent":
			payload := domain.CreateEventPayload{}
			err := UnmarshalPayload(act.Payload, &payload)
			if err != nil {
				return err
			}
			fmt.Printf("- Create %s in %s\n", payload.RefType, act.Repo.Name)
		case "DeleteEvent":
			payload := domain.DeleteEventPayload{}
			err := UnmarshalPayload(act.Payload, &payload)
			if err != nil {
				return err
			}
			fmt.Printf("- Delete %s %s\n", payload.RefType, act.Repo.Name)
		case "ForkEvent":
			fmt.Printf("- Fork repository %s\n", act.Repo.Name)
		case "GollumEvent":
			fmt.Printf("- Modify %s wikis\n", act.Repo.Name)
		case "IssueCommentEvent":
			fmt.Printf("- Comment on issue %s\n", act.Repo.Name)
		case "IssuesEvent":
			fmt.Printf("- Opened a new issue on %s\n", act.Repo.Name)
		case "MemberEvent":
			fmt.Printf("- Member added to %s\n", act.Repo.Name)
		case "PublicEvent":
			fmt.Printf("- Made %s public\n", act.Repo.Name)
		case "PullRequestEvent":
			fmt.Printf("- Pull request activity on %s\n", act.Repo.Name)
		case "PullRequestReviewCommentEvent":
			fmt.Printf("-Comment on a pull request review in %s\n", act.Repo.Name)
		case "PushEvent":
			payload := domain.PushEventPayload{}
			err := UnmarshalPayload(act.Payload, &payload)
			if err != nil {
				return err
			}
			fmt.Printf("- Pushed %d commits to %s\n", payload.Size, act.Repo.Name)
		case "ReleaseEvent":
			fmt.Printf("- Release created in %s\n", act.Repo.Name)
		case "SponsorshipEvent":
			fmt.Printf("- Sponsorship activity for %s\n", act.Repo.Name)
		case "WatchEvent":
			fmt.Printf("- Starred %s\n", act.Repo.Name)
		}
	}

	return nil
}

