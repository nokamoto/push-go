package nokamoto.github.com.pushgoandroid;

import android.os.AsyncTask;
import android.util.Log;

import com.google.firebase.messaging.FirebaseMessagingService;
import com.google.firebase.messaging.RemoteMessage;

import io.grpc.ManagedChannel;
import push.Service;

public class MyFirebaseMessagingService extends FirebaseMessagingService {
    private static final String TAG = "MyFirebaseMsgService";

    /**
     * Called when message is received.
     *
     * @param remoteMessage Object representing the message received from Firebase Cloud Messaging.
     */
    @Override
    public void onMessageReceived(RemoteMessage remoteMessage) {
        // Not getting messages here? See why this may be: https://goo.gl/39bRNJ
        Log.d(TAG, "From: " + remoteMessage.getFrom());

        // Check if message contains a data payload.
        if (remoteMessage.getData().size() > 0) {
            Log.d(TAG, "Message data payload: " + remoteMessage.getData());
        }

        // Check if message contains a notification payload.
        if (remoteMessage.getNotification() != null) {
            Log.d(TAG, "Message Notification Body: " + remoteMessage.getNotification().getBody());
        }
    }

    private class LogTask extends AsyncTask<String, Integer, Integer> {
        @Override
        protected Integer doInBackground(String... logs) {
            ManagedChannel channel = GrpcUtils.newLogChannel();

            try {
                push.LogServiceGrpc.newBlockingStub(channel).info(Service.Log.newBuilder().setText(logs[0]).build());
            } finally {
                channel.shutdown();
            }

            return null;
        }
    }
}
