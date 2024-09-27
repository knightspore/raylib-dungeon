#version 410 core

in vec2 fragTexCoord;
in vec4 fragColor;

out vec4 finalColor;

uniform sampler2D tex;

int size = 3; // blur size
float separation = 3; // spread
float threshold = 0.8;
float amount = 1;

void main() {
    vec2 texSize = textureSize(tex, 0).xy;

    float dist = length(fragTexCoord.y- 0.5);

    float value = 0.0;
    float count = 0.0;

    vec4 result = vec4(0.0);
    vec4 color = vec4(0.0);

    for (int x = -size; x <= size; x++) {
        for (int y = -size; y <= size; y++) {
            color = texture(tex, fragTexCoord + (vec2(x, y) * separation * dist) / texSize);
            value = max(color.r, max(color.g, color.b));
            if (value > threshold) {
                color = vec4(0.0);
            }
            result += color;
            count++;
        }
    }

    finalColor = fragColor * mix(vec4(1.0), result / count, amount);
}
