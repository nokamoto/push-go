package nokamoto.github.com.pushgoandroid;

import android.os.AsyncTask;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.widget.TextView;

import com.google.firebase.iid.FirebaseInstanceId;

import io.grpc.ManagedChannel;
import push.Service;

public class MainActivity extends AppCompatActivity {
    final static String TAG = MainActivity.class.getSimpleName();

    ManagedChannel channel;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        channel = GrpcUtils.newChannel();
    }

    public void register(View view) {
        TextView token = (TextView)findViewById(R.id.token);
        String s = FirebaseInstanceId.getInstance().getToken();

        if (s != null) {
            token.setText(FirebaseInstanceId.getInstance().getToken());
            RegisterTask task = new RegisterTask();
            task.execute(Service.FirebaseCloudMessagingEndpoint.newBuilder().setToken(s).build());
        }
    }

    private class RegisterTask extends AsyncTask<Service.FirebaseCloudMessagingEndpoint, Integer, Integer> {
        @Override
        protected Integer doInBackground(Service.FirebaseCloudMessagingEndpoint... firebaseCloudMessagingEndpoints) {
            Service.SetFirebaseCloudMessagingEndpoint req =
                    Service.SetFirebaseCloudMessagingEndpoint.
                            newBuilder().
                            setKey(Service.Id.newBuilder().setId("x")).
                            setValue(firebaseCloudMessagingEndpoints[0]).
                            build();

            Log.d(TAG, "set: " + req);

            push.FirebaseCloudMessagingEndpointServiceGrpc.newBlockingStub(channel).set(req);

            return null;
        }
    }
}
