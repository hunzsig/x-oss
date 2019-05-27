package scope

import (
	"../response"
	"github.com/kataras/iris"
	"io"
	"os"
)

func Homepage(ctx iris.Context) bool {
	return response.Success(ctx, "Welcome golang", nil)
}

func Upload(ctx iris.Context) bool {
	// Get the file from the request.
	file, header, err := ctx.FormFile("file")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return response.Error(ctx, "Error while uploading: "+err.Error(), nil)
	}
	defer file.Close()
	filename := header.Filename
	fileSize := header.Size
	out, err := os.OpenFile("./uploads/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return response.Error(ctx, "Error while uploading: "+err.Error(), nil)
	}
	defer out.Close()
	io.Copy(out, file)
	return response.Success(ctx, "Upload over"+string(fileSize), nil)
}
