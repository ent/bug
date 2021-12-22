# Template Repository For Bug Reporting
Create a Repository From the Template
-----------------------------------
First and foremost, create a repository from this template repository (the one you're in now! 😊 ).

In order to do so, click 'Use this template' and create your repository.

<img title="use_template" src="https://user-images.githubusercontent.com/73284641/137712207-684ede15-2113-46f1-a214-461976cf10b3.png" width="90%">

Upload your desired code into the new repository.

If you happen to be a GitHub Teams or GitHub Enterprise customer you can alternatively choose to do so directly in the Codespace. Proceed to read if you want to learn how.

Once creating the desired repository, press `.` (period key) from the main repository page to switch over to dev mode. 

Create a Codespace
------------------
Within dev mode, by clicking the 'Run and Debug' button you can create a new Codespace. 

<img title="create_codespace" src="https://user-images.githubusercontent.com/73284641/137910777-3c5e6d79-aefb-49e3-83c0-a180a298881a.png" width="90%">

If you haven't already done so, you can import your code from your personal files by clicking File > Open File. 

<img title="open_file" src="https://user-images.githubusercontent.com/73284641/137134989-2787b722-1d14-4795-8daf-3814289e2e55.png" width="90%">

Inside Codespaces
------------------
Once you're inside Codespaces with your desired code, you can edit, run and debug it.

**Editing**

Let's say you want to edit the template `ent/schema`. For example, add a field named `title` to the `User` schema.

In `ent/schema/user.go`:  

<img title="edit_user" src="https://user-images.githubusercontent.com/73284641/138644354-7c0c51e1-cfb6-4b20-9e40-c275cea352dc.png" width="90%">

To generate this, we will run the command `go generate ./ent` in the Codespaces terminal. 

To see if the command worked successfully, we can see that changes were made to files (files on the left-hand side became yellow, meaning changes were made that were not staged yet). Also, we can check inside `ent/user/user.go` if our field was added: 

<img title = "user changes" src="https://user-images.githubusercontent.com/73284641/138584753-270bf1a7-ab8e-4daa-b71b-b77ab417013c.png" width="90%">

To learn more about working with `ent`, click [here](https://entgo.io/docs/getting-started/).  

**Saving Changes** 

If you have made changes to your code, by clicking on 'Source Control' you can directly stage, commit and push them into the repository.

By clicking the '+' you can stage your changes.

<img title = "stage_changes" src="https://user-images.githubusercontent.com/73284641/137622693-65df9b71-4a07-4eda-9efb-ed882e55e98e.png" width="90%">

To commit the changes, enter a message and click the check mark. This button will automatically commit straight into the master branch.

<img title = "commit_changes" src = "https://user-images.githubusercontent.com/73284641/137622579-603a338b-5dd8-45d4-8922-08ea74bb2245.png" width="90%">


Report a Bug
------------------
Now that your repository is ready, click [here](https://github.com/ent/ent/issues/new/choose) to submit an issue on the Ent repository.   

When submitting your issue, a best practice would be to provide sufficient context, link to your new repository, and you're done!

