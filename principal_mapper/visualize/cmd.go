package visualize

type Visualize struct {
	FileType     string `kong:"name='file-type',default='svg',enum='svg,png,dot,graphml',help='The (lowercase) filetype to output the image as.'"`
	OnlyPrivesc  bool   `kong:"name='only-privesc',help='Generates an image file representing an AWS account.'"`
	WithServices bool   `kong:"name='with-services',help='Includes services with access to Roles in the AWS account visualization.'"`
}
