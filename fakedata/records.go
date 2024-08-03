package fakedata

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/security"

	"github.com/SecurityBrewery/catalyst/migrations"
)

const (
	minimumUserCount   = 1
	minimumTicketCount = 1
)

func Generate(app core.App, userCount, ticketCount int) error {
	if userCount < minimumUserCount {
		userCount = minimumUserCount
	}

	if ticketCount < minimumTicketCount {
		ticketCount = minimumTicketCount
	}

	records, err := Records(app, userCount, ticketCount)
	if err != nil {
		return err
	}

	for _, record := range records {
		if err := app.Dao().SaveRecord(record); err != nil {
			return err
		}
	}

	return nil
}

func Records(app core.App, userCount int, ticketCount int) ([]*models.Record, error) {
	types, err := app.Dao().FindRecordsByExpr(migrations.TypeCollectionName)
	if err != nil {
		return nil, err
	}

	users := userRecords(app.Dao(), userCount)
	tickets := ticketRecords(app.Dao(), users, types, ticketCount)
	reactions := reactionRecords(app.Dao())

	var records []*models.Record
	records = append(records, users...)
	records = append(records, types...)
	records = append(records, tickets...)
	records = append(records, reactions...)

	return records, nil
}

func userRecords(dao *daos.Dao, count int) []*models.Record {
	collection, err := dao.FindCollectionByNameOrId(migrations.UserCollectionName)
	if err != nil {
		panic(err)
	}

	records := make([]*models.Record, 0, count)

	// create the test user
	if _, err := dao.FindRecordById(migrations.UserCollectionName, "u_test"); err != nil {
		record := models.NewRecord(collection)
		record.SetId("u_test")
		_ = record.SetUsername("u_test")
		_ = record.SetPassword("1234567890")
		record.Set("name", gofakeit.Name())
		record.Set("email", "user@catalyst-soar.com")
		_ = record.SetVerified(true)

		records = append(records, record)
	}

	for range count - 1 {
		record := models.NewRecord(collection)
		record.SetId("u_" + security.PseudorandomString(10))
		_ = record.SetUsername("u_" + security.RandomStringWithAlphabet(5, "123456789"))
		_ = record.SetPassword("1234567890")
		record.Set("name", gofakeit.Name())
		record.Set("email", gofakeit.Username()+"@catalyst-soar.com")
		_ = record.SetVerified(true)

		records = append(records, record)
	}

	return records
}

func ticketRecords(dao *daos.Dao, users, types []*models.Record, count int) []*models.Record {
	collection, err := dao.FindCollectionByNameOrId(migrations.TicketCollectionName)
	if err != nil {
		panic(err)
	}

	records := make([]*models.Record, 0, count)

	created := time.Now()
	number := gofakeit.Number(200*count, 300*count)

	for range count {
		number -= gofakeit.Number(100, 200)
		created = created.Add(time.Duration(-gofakeit.Number(13, 37)) * time.Hour)

		record := models.NewRecord(collection)
		record.SetId("t_" + security.PseudorandomString(10))

		updated := gofakeit.DateRange(created, time.Now())

		ticketType := random(types)

		record.Set("created", created.Format("2006-01-02T15:04:05Z"))
		record.Set("updated", updated.Format("2006-01-02T15:04:05Z"))

		record.Set("name", fmt.Sprintf("%s-%d", strings.ToUpper(ticketType.GetString("singular")), number))
		record.Set("type", ticketType.GetId())
		record.Set("description", fakeTicketDescription())
		record.Set("open", gofakeit.Bool())
		record.Set("schema", `{"type":"object","properties":{"tlp":{"title":"TLP","type":"string"}}}`)
		record.Set("state", `{"severity":"Medium"}`)
		record.Set("owner", random(users).GetId())

		records = append(records, record)

		// Add comments
		records = append(records, commentRecords(dao, users, created, record)...)
		records = append(records, timelineRecords(dao, created, record)...)
		records = append(records, taskRecords(dao, users, created, record)...)
		records = append(records, linkRecords(dao, created, record)...)
	}

	return records
}

func commentRecords(dao *daos.Dao, users []*models.Record, created time.Time, record *models.Record) []*models.Record {
	commentCollection, err := dao.FindCollectionByNameOrId(migrations.CommentCollectionName)
	if err != nil {
		panic(err)
	}

	records := make([]*models.Record, 0, 5)

	for range gofakeit.IntN(5) {
		commentCreated := gofakeit.DateRange(created, time.Now())
		commentUpdated := gofakeit.DateRange(commentCreated, time.Now())

		commentRecord := models.NewRecord(commentCollection)
		commentRecord.SetId("c_" + security.PseudorandomString(10))
		commentRecord.Set("created", commentCreated.Format("2006-01-02T15:04:05Z"))
		commentRecord.Set("updated", commentUpdated.Format("2006-01-02T15:04:05Z"))
		commentRecord.Set("ticket", record.GetId())
		commentRecord.Set("author", random(users).GetId())
		commentRecord.Set("message", fakeTicketComment())

		records = append(records, commentRecord)
	}

	return records
}

func timelineRecords(dao *daos.Dao, created time.Time, record *models.Record) []*models.Record {
	timelineCollection, err := dao.FindCollectionByNameOrId(migrations.TimelineCollectionName)
	if err != nil {
		panic(err)
	}

	records := make([]*models.Record, 0, 5)

	for range gofakeit.IntN(5) {
		timelineCreated := gofakeit.DateRange(created, time.Now())
		timelineUpdated := gofakeit.DateRange(timelineCreated, time.Now())

		timelineRecord := models.NewRecord(timelineCollection)
		timelineRecord.SetId("tl_" + security.PseudorandomString(10))
		timelineRecord.Set("created", timelineCreated.Format("2006-01-02T15:04:05Z"))
		timelineRecord.Set("updated", timelineUpdated.Format("2006-01-02T15:04:05Z"))
		timelineRecord.Set("ticket", record.GetId())
		timelineRecord.Set("time", gofakeit.DateRange(created, time.Now()).Format("2006-01-02T15:04:05Z"))
		timelineRecord.Set("message", fakeTicketTimelineMessage())

		records = append(records, timelineRecord)
	}

	return records
}

func taskRecords(dao *daos.Dao, users []*models.Record, created time.Time, record *models.Record) []*models.Record {
	taskCollection, err := dao.FindCollectionByNameOrId(migrations.TaskCollectionName)
	if err != nil {
		panic(err)
	}

	records := make([]*models.Record, 0, 5)

	for range gofakeit.IntN(5) {
		taskCreated := gofakeit.DateRange(created, time.Now())
		taskUpdated := gofakeit.DateRange(taskCreated, time.Now())

		taskRecord := models.NewRecord(taskCollection)
		taskRecord.SetId("ts_" + security.PseudorandomString(10))
		taskRecord.Set("created", taskCreated.Format("2006-01-02T15:04:05Z"))
		taskRecord.Set("updated", taskUpdated.Format("2006-01-02T15:04:05Z"))
		taskRecord.Set("ticket", record.GetId())
		taskRecord.Set("name", fakeTicketTask())
		taskRecord.Set("open", gofakeit.Bool())
		taskRecord.Set("owner", random(users).GetId())

		records = append(records, taskRecord)
	}

	return records
}

func linkRecords(dao *daos.Dao, created time.Time, record *models.Record) []*models.Record {
	linkCollection, err := dao.FindCollectionByNameOrId(migrations.LinkCollectionName)
	if err != nil {
		panic(err)
	}

	records := make([]*models.Record, 0, 5)

	for range gofakeit.IntN(5) {
		linkCreated := gofakeit.DateRange(created, time.Now())
		linkUpdated := gofakeit.DateRange(linkCreated, time.Now())

		linkRecord := models.NewRecord(linkCollection)
		linkRecord.SetId("l_" + security.PseudorandomString(10))
		linkRecord.Set("created", linkCreated.Format("2006-01-02T15:04:05Z"))
		linkRecord.Set("updated", linkUpdated.Format("2006-01-02T15:04:05Z"))
		linkRecord.Set("ticket", record.GetId())
		linkRecord.Set("url", gofakeit.URL())
		linkRecord.Set("name", random([]string{"Blog", "Forum", "Wiki", "Documentation"}))

		records = append(records, linkRecord)
	}

	return records
}

const alertIngestPy = `import sys
import json
import random
import os

from pocketbase import PocketBase

# Parse the event from the webhook payload
event = json.loads(sys.argv[1])
body = json.loads(event["body"])

# Connect to the PocketBase server
client = PocketBase(os.environ["CATALYST_APP_URL"])
client.auth_store.save(token=os.environ["CATALYST_TOKEN"])

# Create a new ticket
client.collection("tickets").create({
	"name": body["name"],
	"type": "alert",
	"open": True,
})`

const assignTicketsPy = `import sys
import json
import random
import os

from pocketbase import PocketBase

# Parse the ticket from the input
ticket = json.loads(sys.argv[1])

# Connect to the PocketBase server
client = PocketBase(os.environ["CATALYST_APP_URL"])
client.auth_store.save(token=os.environ["CATALYST_TOKEN"])

# Get a random user
users = client.collection("users").get_list(1, 200)
random_user = random.choice(users.items)

# Assign the ticket to the random user
client.collection("tickets").update(ticket["record"]["id"], {
	"owner": random_user.id,
})`

const (
	triggerWebhook = `{"token":"1234567890","path":"webhook"}`
	triggerHook    = `{"collections":["tickets"],"events":["create"]}`
)

func reactionRecords(dao *daos.Dao) []*models.Record {
	var records []*models.Record

	collection, err := dao.FindCollectionByNameOrId(migrations.ReactionCollectionName)
	if err != nil {
		panic(err)
	}

	alertIngestActionData, err := json.Marshal(map[string]interface{}{
		"requirements": "pocketbase",
		"script":       alertIngestPy,
	})
	if err != nil {
		panic(err)
	}

	record := models.NewRecord(collection)
	record.SetId("w_" + security.PseudorandomString(10))
	record.Set("name", "Test Reaction")
	record.Set("trigger", "webhook")
	record.Set("triggerdata", triggerWebhook)
	record.Set("action", "python")
	record.Set("actiondata", string(alertIngestActionData))

	records = append(records, record)

	assignTicketsActionData, err := json.Marshal(map[string]interface{}{
		"requirements": "pocketbase",
		"script":       assignTicketsPy,
	})
	if err != nil {
		panic(err)
	}

	record = models.NewRecord(collection)
	record.SetId("w_" + security.PseudorandomString(10))
	record.Set("name", "Test Reaction 2")
	record.Set("trigger", "hook")
	record.Set("triggerdata", triggerHook)
	record.Set("action", "python")
	record.Set("actiondata", string(assignTicketsActionData))

	records = append(records, record)

	return records
}
