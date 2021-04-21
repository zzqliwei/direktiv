package workflow

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/vorteil/direktiv/pkg/cli/util"
)

type workflowObject struct {
	UID       string `json:"uid"`
	ID        string `json:"id"`
	Revision  int    `json:"revision"`
	Active    bool   `json:"active"`
	Createdat struct {
		Seconds int `json:"seconds"`
		Nanos   int `json:"nanos"`
	} `json:"createdAt"`
}

// CreateCommand create the namespace command and subcommands
func CreateCommand() *cobra.Command {

	cmd := util.GenerateCmd("workflows", "List, create, get and execute workflows", "", nil, nil)

	cmd.AddCommand(workflowAddCmd)
	cmd.AddCommand(workflowUpdateCmd)
	cmd.AddCommand(workflowDeleteCmd)
	cmd.AddCommand(workflowListCmd)
	cmd.AddCommand(workflowGetCmd)
	cmd.AddCommand(workflowExecuteCmd)
	cmd.AddCommand(workflowToggleCmd)

	return cmd

}

// Workflow ..
// s.routes[http.MethodGet]["/api/namespaces/{namespace}/workflows/"] = http.HandlerFunc(s.handler.Workflows)
// s.routes[http.MethodGet]["/api/namespaces/{namespace}/workflows/{workflowTarget}"] = http.HandlerFunc(s.handler.GetWorkflow)
// s.routes[http.MethodPut]["/api/namespaces/{namespace}/workflows/{workflowUID}"] = http.HandlerFunc(s.handler.UpdateWorkflow)
// s.routes[http.MethodPut]["/api/namespaces/{namespace}/workflows/{workflowUID}/toggle"] = http.HandlerFunc(s.handler.ToggleWorkflow)
// s.routes[http.MethodPost]["/api/namespaces/{namespace}/workflows"] = http.HandlerFunc(s.handler.CreateWorkflow)
// s.routes[http.MethodDelete]["/api/namespaces/{namespace}/workflows/{workflowUID}"] = http.HandlerFunc(s.handler.DeleteWorkflow)
// s.routes[http.MethodGet]["/api/namespaces/{namespace}/workflows/{workflowUID}/download"] = http.HandlerFunc(s.handler.DownloadWorkflow)
// s.routes[http.MethodPost]["/api/namespaces/{namespace}/workflows/{workflowID}/execute"] = http.HandlerFunc(s.handler.ExecuteWorkflow)

var workflowGetCmd = util.GenerateCmd("get NAMESPACE ID", "Get YAML of a workflow", "", func(cmd *cobra.Command, args []string) {

}, cobra.ExactArgs(2))

var workflowToggleCmd = util.GenerateCmd("toggle NAMESPACE WORKFLOW", "Enables or disables the workflow provided", "", func(cmd *cobra.Command, args []string) {

}, cobra.ExactArgs(2))

var workflowAddCmd = util.GenerateCmd("create NAMESPACE WORKFLOW FILE", "Creates a new workflow on provided namespace", "", func(cmd *cobra.Command, args []string) {

	// read file
	f, err := ioutil.ReadFile(args[1])
	if err != nil {
		log.Fatalf("can not create workflow: %v", err)
	}
	st := string(f)

	wf, err := util.DoRequest(http.MethodPost, fmt.Sprintf("/namespaces/%s/workflows",
		args[0]), util.YAMLCt, &st)
	if err != nil {
		log.Fatalf("error creating workflow: %v", err)
	}

	var workflow workflowObject
	err = json.Unmarshal(wf, &workflow)
	if err != nil {
		log.Fatalf("can not parse response: %v, %v", err, string(wf))
	}

	log.Printf("workflow '%s' created", workflow.ID)

}, cobra.ExactArgs(2))

var workflowUpdateCmd = util.GenerateCmd("update NAMESPACE NAME FILE", "Updates an existing workflow", "", func(cmd *cobra.Command, args []string) {

	// read file
	f, err := ioutil.ReadFile(args[2])
	if err != nil {
		log.Fatalf("can not create workflow: %v", err)
	}
	st := string(f)

	wf, err := util.DoRequest(http.MethodPut, fmt.Sprintf("/namespaces/%s/workflows/%s",
		args[0], args[1]), util.YAMLCt, &st)
	if err != nil {
		log.Fatalf("error updating workflow: %v", err)
	}

	var workflow workflowObject
	err = json.Unmarshal(wf, &workflow)
	if err != nil {
		log.Fatalf("can not parse response: %v, %v", err, string(wf))
	}

	log.Printf("workflow '%s' updated", workflow.ID)

}, cobra.ExactArgs(3))

var workflowDeleteCmd = util.GenerateCmd("delete NAMESPACE NAME", "Deletes an existing workflow", "", func(cmd *cobra.Command, args []string) {

}, cobra.ExactArgs(2))

var workflowListCmd = util.GenerateCmd("list NAMESPACE", "List all workflows under a namespace", "", func(cmd *cobra.Command, args []string) {

	// /api/namespaces/{namespace}/workflows/

}, cobra.ExactArgs(1))

// workflowExecuteCmd
var workflowExecuteCmd = util.GenerateCmd("execute NAMESPACE ID", "Executes workflow with provided ID", "", func(cmd *cobra.Command, args []string) {

}, cobra.ExactArgs(2))

// Toggle enables or disables the workflow
// func Toggle(conn *grpc.ClientConn, namespace, workflow string) (string, error) {
// 	client, ctx, cancel := util.CreateClient(conn)
// 	defer cancel()
//
// 	request := ingress.GetWorkflowByIdRequest{
// 		Namespace: &namespace,
// 		Id:        &workflow,
// 	}
//
// 	resp, err := client.GetWorkflowById(ctx, &request)
// 	if err != nil {
// 		s := status.Convert(err)
// 		return "", fmt.Errorf("[%v] %v", s.Code(), s.Message())
// 	}
//
// 	toggle := !*resp.Active
//
// 	uRequest := ingress.UpdateWorkflowRequest{
// 		Uid:      resp.Uid,
// 		Workflow: resp.Workflow,
// 		Active:   &toggle,
// 	}
//
// 	_, err = client.UpdateWorkflow(ctx, &uRequest)
// 	if err != nil {
// 		s := status.Convert(err)
// 		return "", fmt.Errorf("[%v] %v", s.Code(), s.Message())
// 	}
//
// 	if toggle {
// 		return fmt.Sprintf("Enabled workflow '%s'", workflow), nil
// 	}
//
// 	return fmt.Sprintf("Disabled workflow '%s'", workflow), nil
// }
//
// // List returns an array of workflows for a given namespace
// func List(conn *grpc.ClientConn, namespace string) ([]*ingress.GetWorkflowsResponse_Workflow, error) {
// 	client, ctx, cancel := util.CreateClient(conn)
// 	defer cancel()
//
// 	// prepare request
// 	request := ingress.GetWorkflowsRequest{
// 		Namespace: &namespace,
// 	}
//
// 	// send grpc request
// 	resp, err := client.GetWorkflows(ctx, &request)
// 	if err != nil {
// 		s := status.Convert(err)
// 		return nil, fmt.Errorf("[%v] %v", s.Code(), s.Message())
// 	}
//
// 	return resp.Workflows, nil
// }
//
// // Execute a workflow using the yaml provided
// func Execute(conn *grpc.ClientConn, namespace string, id string, input string) (string, error) {
// 	client, ctx, cancel := util.CreateClient(conn)
// 	defer cancel()
//
// 	var err error
// 	var b []byte
// 	if input != "" {
// 		b, err = ioutil.ReadFile(input)
// 		if err != nil {
// 			return "", err
// 		}
// 	}
//
// 	// prepare request
// 	request := ingress.InvokeWorkflowRequest{
// 		Namespace:  &namespace,
// 		Input:      b,
// 		WorkflowId: &id,
// 	}
//
// 	// send grpc request
// 	resp, err := client.InvokeWorkflow(ctx, &request)
// 	if err != nil {
// 		s := status.Convert(err)
// 		return "", fmt.Errorf("[%v] %v", s.Code(), s.Message())
// 	}
//
// 	return fmt.Sprintf("Successfully invoked, Instance ID: %s", resp.GetInstanceId()), nil
// }
//
// // getWorkflowUID returns the UID of a workflow
// func getWorkflowUID(ctx context.Context, client ingress.DirektivIngressClient, namespace, id string) (string, error) {
// 	// defer cancel()
// 	// prepare request
// 	request := ingress.GetWorkflowByIdRequest{
// 		Namespace: &namespace,
// 		Id:        &id,
// 	}
//
// 	// send grpc request
// 	resp, err := client.GetWorkflowById(ctx, &request)
// 	if err != nil {
// 		s := status.Convert(err)
// 		return "", fmt.Errorf("[%v] %v", s.Code(), s.Message())
// 	}
// 	return resp.GetUid(), nil
// }
//
// // Get returns a workflow definition in YAML format
// func Get(conn *grpc.ClientConn, namespace string, id string) (string, error) {
// 	client, ctx, cancel := util.CreateClient(conn)
// 	defer cancel()
//
// 	// prepare request
// 	request := ingress.GetWorkflowByIdRequest{
// 		Namespace: &namespace,
// 		Id:        &id,
// 	}
//
// 	// send grpc request
// 	resp, err := client.GetWorkflowById(ctx, &request)
// 	if err != nil {
// 		s := status.Convert(err)
// 		return "", fmt.Errorf("[%v] %v", s.Code(), s.Message())
// 	}
//
// 	return string(resp.GetWorkflow()), nil
// }
//
// // Update a workflow specified by ID.
// func Update(conn *grpc.ClientConn, namespace string, id string, filepath string) (string, error) {
// 	client, ctx, cancel := util.CreateClient(conn)
// 	defer cancel()
//
// 	b, err := ioutil.ReadFile(filepath)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	uid, err := getWorkflowUID(ctx, client, namespace, id)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	// prepare request
// 	request := ingress.UpdateWorkflowRequest{
// 		Uid:      &uid,
// 		Workflow: b,
// 	}
//
// 	// send grpc request
// 	resp, err := client.UpdateWorkflow(ctx, &request)
// 	if err != nil {
// 		s := status.Convert(err)
// 		return "", fmt.Errorf("[%v] %v", s.Code(), s.Message())
// 	}
//
// 	return fmt.Sprintf("Successfully updated '%s'", resp.GetId()), nil
// }
//
// // Delete an existing workflow.
// func Delete(conn *grpc.ClientConn, namespace, id string) (string, error) {
// 	client, ctx, cancel := util.CreateClient(conn)
// 	defer cancel()
//
// 	uid, err := getWorkflowUID(ctx, client, namespace, id)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	// prepare request
// 	request := ingress.DeleteWorkflowRequest{
// 		Uid: &uid,
// 	}
//
// 	// send grpc request
// 	_, err = client.DeleteWorkflow(ctx, &request)
// 	if err != nil {
// 		s := status.Convert(err)
// 		return "", fmt.Errorf("[%v] %v", s.Code(), s.Message())
// 	}
//
// 	return fmt.Sprintf("Deleted workflow '%v'", id), nil
// }
//
// // Add creates a new workflow
// func Add(conn *grpc.ClientConn, namespace string, filepath string) (string, error) {
// 	client, ctx, cancel := util.CreateClient(conn)
// 	defer cancel()
//
// 	b, err := ioutil.ReadFile(filepath)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	active := true
//
// 	// prepare request
// 	request := ingress.AddWorkflowRequest{
// 		Namespace: &namespace,
// 		Workflow:  b,
// 		Active:    &active,
// 	}
//
// 	// send grpc request
// 	resp, err := client.AddWorkflow(ctx, &request)
// 	if err != nil {
// 		s := status.Convert(err)
// 		return "", fmt.Errorf("[%v] %v", s.Code(), s.Message())
// 	}
//
// 	return fmt.Sprintf("Created workflow '%s'", resp.GetId()), nil
// }
